import { render } from '@testing-library/react';
import { Health } from '../../interfaces/Index';

import Home from '../../pages/index';

const health: Health[] = [
  {
    agentID: '1',
    createTime: 1,
    updateTime: 1,
    id: '1',
    online: true,
    host: {
      os: 'ubunut',
      platform: 'linux',
      hostname: 'Test Machine',
    },
  },
];

test('Index page matches snapshot', () => {
  const { asFragment } = render(<Home error="" health={health} />);
  expect(asFragment()).toMatchSnapshot();
});
