import { GetServerSideProps, InferGetServerSidePropsType } from 'next';
import { useEffect, useState } from 'react';

import AgentStats from '../components/Headers/AgentStats';
import AgentInformation from '../components/Tables/AgentInformation';
import { Health, Host, HostResponse } from '../interfaces/Index';
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

  useEffect(() => {
    const ws = new WebSocket('wss://localhost:3000/ws/v1/health/');

    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      const now = new Date().valueOf();

      // TODO:
      // Loop through hosts
      // Get all health for that host
      // Get last health packet (since sorted)
      // Update online/lastConnected
      data.Data.forEach((health: Health) => {
        hosts?.forEach((host) => {
          if (host.agentID == health.agentID) {
            if (health.createTime > (host.lastConnected || 0)) {
              host.lastConnected = health.createTime;
              // TODO: move this to a utils file
              host.online =
                parseInt((host.lastConnected || 0).toString().substr(0, 13)) >
                now;
            }
            host.health = data.Data;
            return;
          }
        });
      });
      setHosts(hosts);
      console.log(hosts);
    };

    return () => {
      ws.close();
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
