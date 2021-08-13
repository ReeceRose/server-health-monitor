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

  const [hosts, setHosts] = useState(initial_hosts);
  const websocket = useRef<WebSocket | null>(null);

  useEffect(() => {
    websocket.current = new WebSocket('wss://localhost:3000/ws/v1/health/');

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
        // let updateThisHost = false;
        if (host.health !== null) {
          if (host.health === undefined)
            //|| host.health[0].createTime === 0
            return;
          hosts[index].health?.push(...host.health);
          updateHosts = true;
          const latestHealth = host.health[0];
          if (
            latestHealth.createTime >
            (host.lastConnected || hosts[index].lastConnected || 0)
          ) {
            hosts[index].lastConnected = latestHealth.createTime;
            updateHosts = true;
            // updateThisHost = true;
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
          (parseInt(process.env.HEALTH_DELAY || '5') || 5) * 60 * 1000;
        const online =
          now.valueOf() - lastConnected.valueOf() < maximumTimeSinceLastConnect;
        if (online != hosts[index].online) {
          hosts[index].online = online;
          updateHosts = true;
          // updateThisHost = true;
          if (host.lastConnected === 0) {
            host.lastConnected = hosts[index].lastConnected;
          }
        }

        // if (updateThisHost) {
        //   hosts[index] = host;
        // }
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
