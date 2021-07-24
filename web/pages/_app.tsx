import { AppProps } from 'next/app';
import { useEffect } from 'react';
import SEO from '../components/SEO';
import Layout from '../layouts/Default';

import '../styles/globals.css';
import '@fortawesome/fontawesome-free/css/all.min.css';

function MyApp({ Component, pageProps }: AppProps): JSX.Element {
  useEffect(() => {
    // Note, a button is required to switch the themes and it will also need to call classList.add.
    if (localStorage.theme) {
      document.documentElement.classList.add(localStorage.theme);
    }
  }, []);

  return (
    <>
      <SEO
        title="Server Health Monitor"
        description="Monitor servers health from one central location"
      />
      <div className="flex flex-wrap">
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </div>
    </>
  );
}

export default MyApp;
