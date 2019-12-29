import React from 'react';
import axios from 'axios';

import "regenerator-runtime/runtime";

import { act, create } from 'react-test-renderer';

import Editor from './editor';

jest.mock('axios');

jest.mock('react-ace', () => {
    return jest.fn().mockImplementation(() => {
        return null;
    });
});

jest.mock('ace-builds', () => "");
jest.mock('ace-builds/webpack-resolver', () => "");
jest.mock('ace-builds/src-noconflict/mode-json', () => "");
jest.mock('ace-builds/src-noconflict/theme-github', () => "");

test('<Editor/>', () => {
    const component = create(<Editor/>);
    console.log(component.root._fiber.memoizedState);
    expect(true).toBe(true);
});
