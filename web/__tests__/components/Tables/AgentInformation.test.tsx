import { render } from '@testing-library/react';

import AgentInformation from '../../../components/Tables/AgentInformation';
import { Host } from '../../../interfaces/Index';

const host: Host = {
  os: 'ubunut',
  platform: 'linux',
  hostname: 'Test Machine',
};

test('AgentInformation component matches snapshot', () => {
  const { asFragment } = render(<AgentInformation host={host} online={true} />);
  expect(asFragment()).toMatchSnapshot();
});
