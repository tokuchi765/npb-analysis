import axios, { AxiosResponse } from 'axios';

const rest = (() => {
  const client = axios.create({
    baseURL: 'http://localhost:8081',
    timeout: 15000,
  });
  return {
    client,
    get: <T = any, R = AxiosResponse<T>>(url: string): Promise<R> => {
      return client.get(url);
    },
  };
})();

export { rest };
