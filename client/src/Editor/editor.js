'use strict';

import React, { useState, useEffect, Component } from 'react';

import AceEditor from 'react-ace';

import "ace-builds/webpack-resolver";
import "ace-builds/src-noconflict/mode-json";
import "ace-builds/src-noconflict/theme-github.js";

const Editor = () => {
    const [config, setConfig] = useState("{}");

    const handleChange = value => {
        setConfig(value);
    };

    const handleClick = async e => {
        console.log(e);
    };

    return (
            <div className="h-100">
              <AceEditor
                mode="json"
                theme="github"
                name="json-editor"
                value={config}
                className="mb-2"
                onChange={ handleChange } />
              <button
                type="button"
                onClick={handleClick}
                className="btn btn-primary">
                Encrypt
              </button>
            </div>
    );
};


export default Editor;
