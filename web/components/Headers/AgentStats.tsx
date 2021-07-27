import React from 'react';
import AgentStats from '../Cards/AgentStats';

type Props = {
  active: number;
  inactive: number;
};

const AgentStatsHeader: React.FC<Props> = ({ active, inactive }) => {
  return (
    <div className="relative pt-12 pb-32 bg-gray-800 md:pt-32">
      <div className="w-full px-4 mx-auto md:px-10">
        <div>
          <div className="flex flex-wrap">
            <div className="w-full px-4 lg:w-6/12 xl:w-3/12">
              <AgentStats
                subtitle="ACTIVE SERVERS"
                title={active.toString()}
                iconColour="bg-green-400"
              />
            </div>
            <div className="w-full px-4 lg:w-6/12 xl:w-3/12">
              <AgentStats
                subtitle="INACTIVE SERVERS"
                title={inactive.toString()}
                iconColour="bg-red-400"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default AgentStatsHeader;
