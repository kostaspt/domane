import http from '@store/http';

export function search(query: string) {
  return http.post(`domains/extensions`, { query });
}
