import { GetServerSideProps, InferGetServerSidePropsType } from 'next';
import { useEffect, useRef, useState } from 'react';

import AgentStats from '../components/Headers/AgentStats';
import AgentInformation from '../components/Tables/AgentInformation';
import { Host, HostResponse } from '../interfaces/Index';
import hostService from '../services/host.service';

type Props = {
  initial_hosts: Host[];
  error: string;
};

export const getServerSideProps: GetServerSideProps<Props> = async () => {
  return await hostService.getAll().then((res) => {
    const data: HostResponse = JSON.parse(JSON.stringify(res.data));
    return { props: { initial_hosts: data.Data, error: data.Error } };
  });
};

function Index({
  initial_hosts,
  error,
}: InferGetServerSidePropsType<typeof getServerSideProps>): JSX.Element {
  // TODO: better error handling
  if (error) {
    alert(error);
  }

  const webSocketURL = process.env.WS_URL || 'wss://localhost:3000/ws/v1/';

  const [hosts, setHosts] = useState(initial_hosts);
  const websocket = useRef<WebSocket | null>(null);

  useEffect(() => {
    websocket.current = new WebSocket(webSocketURL + 'health/');

    return () => {
      if (!websocket.current) return;
      websocket.current.close();
    };
  }, []);

  useEffect(() => {
    if (!websocket.current) return;

    websocket.current.onmessage = (event) => {
      const response = JSON.parse(event.data);
      const now = new Date();
      let updateHosts = false;
      response.Data?.forEach((host: Host) => {
        const index = hosts.findIndex((h: Host) => h.agentID == host.agentID);
        if (host.health !== null) {
          if (host.health === undefined) return;
          hosts[index].health?.push(...host.health);
          updateHosts = true;
          const latestHealth = host.health[0];
          if (
            latestHealth.createTime >
            (host.lastConnected || hosts[index].lastConnected || 0)
          ) {
            hosts[index].lastConnected = latestHealth.createTime;
            updateHosts = true;
          }
        }

        const lastConnected = new Date(
          parseInt(
            (host.lastConnected || hosts[index].lastConnected || 0)
              .toString()
              .substr(0, 13)
          )
        );
        const maximumTimeSinceLastConnect =
          (parseInt(process.env.MINUTES_SINCE_HEALTH_SHOW_OFFLINE || '5') ||
            5) *
          60 *
          1000;
        const online =
          now.valueOf() - lastConnected.valueOf() < maximumTimeSinceLastConnect;
        if (online != hosts[index].online) {
          hosts[index].online = online;
          updateHosts = true;

          if (host.lastConnected === 0) {
            host.lastConnected = hosts[index].lastConnected;
          }
        }
      });

      if (updateHosts) {
        setHosts([...hosts]);
      }
    };
  }, [hosts]);

  return (
    <>
      <AgentStats
        active={hosts?.filter((h) => h.online).length || 0}
        inactive={hosts?.filter((h) => !h.online).length || 0}
      />
      <div className="flex-grow w-full px-4 mx-auto mb-0 -m-24 md:px-10">
        <div className="flex flex-wrap">
          <div className="w-full px-4 mb-12 xl:mb-0">
            {hosts?.map((host) => (
              <AgentInformation key={host.hostname} host={host} />
            ))}
          </div>
        </div>
      </div>
    </>
  );
}

export default Index;
