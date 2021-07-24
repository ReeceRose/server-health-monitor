import { render } from '@testing-library/react';

import AgentInformation from '../../../components/Tables/AgentInformation';

test('AgentInformation component matches snapshot', () => {
  const { asFragment } = render(<AgentInformation />);
  expect(asFragment()).toMatchSnapshot();
});
