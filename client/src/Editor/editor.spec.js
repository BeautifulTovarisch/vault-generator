import React from 'react';
import axios from 'axios';

import "regenerator-runtime/runtime";

import { act, create } from 'react-test-renderer';

import Editor from './editor';
import { encryptConfig } from './editor.api';

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

afterEach(() => {
    jest.clearAllMocks();
});

test('<Editor/> default state', async () => {
    const props = {
        setError: jest.fn(),
        setOutput: jest.fn()
    };

    let component;
    act(() => { component = create(<Editor {...props} />); });

    const button = component.root.findByType("button");
    expect(button.props.children).toBe("Encrypt");

    await act(button.props.onClick);

    // Don't call without ensuring encryption key is set
    expect(props.setError).toBeCalled();
    expect(axios.post).not.toBeCalled();
});
