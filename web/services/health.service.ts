import { AxiosPromise } from 'axios';
import client from '../client';

class HealthService {
  getAll(): AxiosPromise {
    return client.get('/api/v1/health/');
  }
}

export default new HealthService();
