import http from '@store/http';

export function search(query) {
  return http.post(`domains/extensions`, { query });
}
