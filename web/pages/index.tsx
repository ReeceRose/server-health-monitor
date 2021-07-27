import { NextPage } from 'next';

import AgentStats from '../components/Headers/AgentStats';
import AgentInformation from '../components/Tables/AgentInformation';
import { Health, HealthResponse } from '../interfaces/Index';
import healthService from '../services/health.service';

type Props = {
  health: Health[];
  error: string;
};

const Index: NextPage<Props> = ({ health, error }) => {
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
            {/* TODO: extract host from health on API side */}
            <AgentInformation online={true} host={health[0].host} />
          </div>
        </div>
      </div>
    </>
  );
};

Index.getInitialProps = async () => {
  return await healthService.getAll().then((res) => {
    const data: HealthResponse = JSON.parse(JSON.stringify(res.data));
    return { health: data.Data, error: data.Error };
  });
};

export default Index;
