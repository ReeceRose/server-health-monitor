import { render } from '@testing-library/react';

import AgentInformation from '../../../components/Tables/AgentInformation';
import { Host } from '../../../interfaces/Index';

const host: Host = {
  os: 'ubunut',
  platform: 'linux',
  hostname: 'Test Machine',
  online: true,
};

test('AgentInformation component matches snapshot', () => {
  const { asFragment } = render(<AgentInformation host={host} />);
  expect(asFragment()).toMatchSnapshot();
});
