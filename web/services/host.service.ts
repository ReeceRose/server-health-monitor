import { AxiosPromise } from 'axios';
import client from '../client';

class HostService {
  getAll(): AxiosPromise {
    return client.get('/api/v1/host/');
  }
}

export default new HostService();
