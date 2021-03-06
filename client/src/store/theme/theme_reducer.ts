import * as types from './types';
import { ThemeAction } from './ThemeActionTypes';
export type Theme = {
  colors: {[key: string]: string}
  common: {
    transition: string;
  }
  borders: {
    radius: string;
    base: string;
  }
  shadows: {
    base: string;
  }
}

const initialState: Theme = {
  colors: {
    brand: '#EEF1F0',
    brandDark: '#111111',
    green: '#14CD73',
    greenDark: '#0CB863',
    red: '#F55466',
    shadow: '#F1F1F1',
    white: '#ffffff',
    black: '#000000',
    bg: '#F3F4F5',
    text: '#DEDEDE',
    textDark: '#333333',
  },
  common: {
    transition: '.2s all ease'
  },
  borders: {
    radius: '5px',
    base: '1px solid #dfdfdf',
  },
  shadows: {
    base: '1px 2px 4px rgba(0,0,0,.03)'
  }
}

export const ThemeReducer = (
  state = initialState,
  action: ThemeAction,
) => {
  switch(action.type) {
    case(types.CHANGE_THEME): {
      return state;
    }
    default: {
      return state;
    }
  }
}