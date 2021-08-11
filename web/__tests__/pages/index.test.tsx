import { render } from '@testing-library/react';
import { Host } from '../../interfaces/Index';

import Index from '../../pages/index';

const hosts: Host[] = [
  {
    hostID: '123',
    os: 'ubunut',
    platform: 'linux',
    hostname: 'Test Machine',
  },
];

test('Index page matches snapshot', () => {
  const { asFragment } = render(<Index error="" initial_hosts={hosts} />);
  expect(asFragment()).toMatchSnapshot();
});
