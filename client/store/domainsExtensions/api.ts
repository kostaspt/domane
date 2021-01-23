import http from '@store/http';

export function apiRequest(query: string) {
  return http.post(`domains/extensions`, { query });
}
