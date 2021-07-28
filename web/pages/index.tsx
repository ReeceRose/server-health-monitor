import { NextPage } from 'next';

import AgentStats from '../components/Headers/AgentStats';
import AgentInformation from '../components/Tables/AgentInformation';
import { Host, HostResponse } from '../interfaces/Index';
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

  return (
    <>
      <AgentStats active={1} inactive={0} />
      <div className="flex-grow w-full px-4 mx-auto mb-0 -m-24 md:px-10">
        <div className="flex flex-wrap">
          <div className="w-full px-4 mb-12 xl:mb-0">
            {hosts.map((host) => (
              <AgentInformation key={host.hostID} host={host} />
            ))}
          </div>
        </div>
      </div>
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
