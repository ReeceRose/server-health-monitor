import { render } from '@testing-library/react';

import AgentStatsHeader from '../../../components/Headers/AgentStats';

test('AgentStats component matches snapshot', () => {
  const { asFragment } = render(<AgentStatsHeader inactive={0} active={1} />);
  expect(asFragment()).toMatchSnapshot();
});
