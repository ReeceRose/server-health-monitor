import Navbar from '../components/Navbar';
import Footer from '../components/Footer';

type Props = {
  children: JSX.Element[] | JSX.Element;
};

const Default: React.FC<Props> = ({ children }) => {
  return (
    <div className="relative flex flex-col w-full min-h-screen bg-gray-100">
      <Navbar />
      {children}
      <Footer />
    </div>
  );
};

export default Default;
