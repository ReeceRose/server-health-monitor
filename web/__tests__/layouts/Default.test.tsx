import { render } from '@testing-library/react';

import Layout from '../../layouts/Default';

test('Default layout page matches snapshot', () => {
  const { asFragment } = render(
    <Layout>
      <></>
    </Layout>
  );
  expect(asFragment()).toMatchSnapshot();
});

test('Default layout page renders children', () => {
  render(
    <Layout>
      <p id="_id">1234</p>
    </Layout>
  );
  expect(document.querySelector('#_id')?.textContent).toBe('1234');
});
