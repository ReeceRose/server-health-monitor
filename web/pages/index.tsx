import { NextPage } from 'next';
import { useEffect, useState } from 'react';

import AgentStats from '../components/Headers/AgentStats';
import AgentInformation from '../components/Tables/AgentInformation';
import { Health, Host, HostResponse } from '../interfaces/Index';
import hostService from '../services/host.service';

type Props = {
  hosts: Host[];
  error: string;
};

const Index: NextPage<Props> = ({ hosts, error }) => {
  // TODO: better error handling
  if (error) {
    alert(error);
  }

  const [health, setHealth] = useState([]);
  // const ws = useRef(WebSocket);

  useEffect(() => {
    const ws = new WebSocket('wss://localhost:3000/ws/v1/health/');

    ws.onmessage = (event) => {
      const parsed_data = JSON.parse(event.data);
      console.log(parsed_data);
      setHealth(parsed_data.Data);
      // TODO: refactor this
      parsed_data.Data.forEach((data: Health) => {
        hosts?.forEach((host) => {
          if (host.hostID == data.agentID) {
            console.log('found match');
            if (data.createTime > (host.lastConnected || 0)) {
              host.lastConnected = data.createTime;
              console.log('host updated');
            }
            return;
          }
        });
      });
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
      {health?.map((health: Health) => (
        <p key={health.createTime}>{health.agentID}</p>
      ))}
    </>
  );
};

Index.getInitialProps = async () => {
  return await hostService.getAll().then((res) => {
    const data: HostResponse = JSON.parse(JSON.stringify(res.data));
    return { hosts: data.Data, error: data.Error };
  });
};

export default Index;
