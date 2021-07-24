import { render } from '@testing-library/react';

import AgentStats from '../../../components/Headers/AgentStats';

test('AgentStats component matches snapshot', () => {
  const { asFragment } = render(<AgentStats />);
  expect(asFragment()).toMatchSnapshot();
});
