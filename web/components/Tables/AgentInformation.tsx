import { Host } from '../../interfaces/Index';

type Props = {
  online: boolean;
  host: Host;
};

const AgentInformation: React.FC<Props> = ({ online, host }) => {
  return (
    <div
      className={
        'relative flex flex-col break-words w-full mb-6 shadow-lg rounded bg-gray-700 text-white'
      }
    >
      <div className="px-4 py-3 mb-0 border-0 rounded-t">
        <div className="flex flex-wrap items-center">
          <div className="relative flex-1 flex-grow w-full max-w-full px-2">
            <h3 className={'font-semibold text-lg text-white'}>Servers</h3>
          </div>
        </div>
      </div>
      <div className="block w-full overflow-x-auto">
        <table className="items-center w-full bg-transparent border-collapse">
          <thead>
            <tr>
              <th
                className={
                  'px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-gray-600 text-gray-200 border-gray-500'
                }
              >
                Hostname
              </th>
              <th
                className={
                  'px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-gray-600 text-gray-200 border-gray-500'
                }
              >
                OS
              </th>
              <th
                className={
                  'px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-gray-600 text-gray-200 border-gray-500'
                }
              >
                Platform
              </th>
              <th
                className={
                  'px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-gray-600 text-gray-200 border-gray-500'
                }
              >
                Last Ping
              </th>
              <th
                className={
                  'px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left bg-gray-600 text-gray-200 border-gray-500'
                }
              >
                Status
              </th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <th className="flex items-center p-4 px-6 text-xs text-left align-middle border-t-0 border-l-0 border-r-0 whitespace-nowrap">
                <span className="text-white">{host.hostname}</span>
              </th>
              <td className="p-4 px-6 text-xs align-middle border-t-0 border-l-0 border-r-0 whitespace-nowrap">
                <span className="text-white">{host.os}</span>
              </td>
              <td className="p-4 px-6 text-xs align-middle border-t-0 border-l-0 border-r-0 whitespace-nowrap">
                <span className="text-white">{host.platform}</span>
              </td>
              <td className="p-4 px-6 text-xs align-middle border-t-0 border-l-0 border-r-0 whitespace-nowrap">
                TODO
              </td>
              <td className="p-4 px-6 text-xs align-middle border-t-0 border-l-0 border-r-0 whitespace-nowrap">
                {online ? (
                  <span>
                    <i className="mr-2 text-green-400 fas fa-circle"></i> Online
                  </span>
                ) : (
                  <span>
                    <i className="mr-2 text-red-400 fas fa-circle"></i> Offline
                  </span>
                )}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default AgentInformation;
