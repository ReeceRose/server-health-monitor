import { render } from '@testing-library/react';

import AgentStats from '../../../components/Cards/AgentStats';

const ToRender = (
  <AgentStats title="2" iconColour="bg-green-500" subtitle="ACTIVE SERVERS" />
);

test('AgentStats component matches snapshot', () => {
  const { asFragment } = render(ToRender);
  expect(asFragment()).toMatchSnapshot();
});

test('AgentStats component data matches passed props', () => {
  render(ToRender);
  expect(document.querySelector('h5')?.textContent).toBe('ACTIVE SERVERS');
  expect(document.querySelector('span')?.textContent).toBe('2');
});
