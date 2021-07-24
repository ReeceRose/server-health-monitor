// import Navbar from '../components/Navbar';
// import StatsHeader from '../components/Headers/Stats';
import Footer from '../components/Footer';

type Props = {
  children: JSX.Element[] | JSX.Element;
};

const Default: React.FC<Props> = ({ children }) => {
  return (
    <div className="relative flex flex-col w-full min-h-screen bg-blueGray-100">
      {/* <Navbar /> */}
      {/* Header */}
      {/* <StatsHeader /> */}
      <div className="flex-grow px-4 mx-auto mb-0 -m-24 md:px-10">
        <div>{children}</div>
      </div>
      <Footer />
    </div>
  );
};

export default Default;
