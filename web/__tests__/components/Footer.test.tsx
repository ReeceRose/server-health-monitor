import { render } from '@testing-library/react';

import Footer from '../../components/Footer';

test('Footer component matches snapshot', () => {
  const { asFragment } = render(<Footer />);
  expect(asFragment()).toMatchSnapshot();
});
