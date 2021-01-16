import { createSlice } from '@reduxjs/toolkit';
import { RootState } from '@store/index';

const name: string = 'query';

export type QueryInitialState = {
  text: string;
};

const initialState: QueryInitialState = {
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

export const selectQueryText = (s: RootState) => s.text;

export const querySliceName = name;
export const { update: updateQuery } = slice.actions;
export default slice.reducer;
