import Navbar from '../components/Navbar';
import AgentStats from '../components/Headers/AgentStats';
import Footer from '../components/Footer';

type Props = {
  children: JSX.Element[] | JSX.Element;
};

const Default: React.FC<Props> = ({ children }) => {
  return (
    <div className="relative flex flex-col w-full min-h-screen bg-gray-100">
      <Navbar />
      <AgentStats />
      <div className="flex-grow w-full px-4 mx-auto mb-0 -m-24 md:px-10">
        {children}
      </div>
      <Footer />
    </div>
  );
};

export default Default;
