import React from 'react';
import { render, fireEvent, waitForElement } from '@testing-library/react';

import Editor from './editor';

jest.mock('react-ace', () => {
    return jest.fn().mockImplementation(() => {
        return null;
    })
});

jest.mock('ace-builds', () => "");
jest.mock('ace-builds/webpack-resolver', () => "");
jest.mock('ace-builds/src-noconflict/mode-json', () => "");
jest.mock('ace-builds/src-noconflict/theme-github', () => "");

test('oi', () => {
    const { getByText } = render(<Editor />);
    expect(true).toBe(true);
})
