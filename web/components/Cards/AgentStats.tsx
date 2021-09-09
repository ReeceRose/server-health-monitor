type Props = {
  title: string;
  subtitle: string;
  iconColour: string;
};

const AgentStatsCard: React.FC<Props> = ({ title, subtitle, iconColour }) => {
  return (
    <div className="relative flex flex-col min-w-0 mb-6 break-words bg-white rounded shadow-lg xl:mb-0">
      <div className="flex-auto p-4">
        <div className="flex flex-wrap">
          <div className="relative flex-1 flex-grow w-full max-w-full pr-4">
            <h5 className="text-xs font-bold text-gray-400 uppercase">
              {subtitle}
            </h5>
            <span className="text-xl font-semibold text-gray-700">{title}</span>
          </div>
          <div className="relative flex-initial w-auto pl-4">
            <div
              className={
                'text-white p-3 text-center inline-flex items-center justify-center w-12 h-12 shadow-lg rounded-full ' +
                iconColour
              }
            >
              <i className="fas fa-server"></i>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default AgentStatsCard;
