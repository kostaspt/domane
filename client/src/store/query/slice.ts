import { createSlice } from '@reduxjs/toolkit';
import { RootState } from '@store';

export const name: string = 'query';

export type State = {
  text: string;
};

const initialState: State = {
  text: '',
};

const slice = createSlice({
  name,
  initialState,
  reducers: {
    update(state, { payload }) {
      state.text = payload;
    },
  },
});

export const selectQueryText = (s: RootState) => (<State>s[name]).text;

export const { update: updateQuery } = slice.actions;
export default slice.reducer;
