'use strict';

import React, { useState } from 'react';

import AceEditor from 'react-ace';

import "ace-builds/webpack-resolver";
import "ace-builds/src-noconflict/mode-json";
import "ace-builds/src-noconflict/theme-github.js";

import { encryptConfig } from './editor.api';

const Editor = ({ setError, setOutput }) => {
    // We don't define a default key here as it should prevent the post request if undefined
    const [state, setState] = useState({
        config: "{}",
        loading: false
    });

    const _handleKeyChange = useCallback(e => {
        e.persist();
        setState(state => ({ ...state, key: e.target.value }));
    });

    const _handleConfigChange = value => {
        setState(state => ({ ...state, config: value }));
    };

    const _handleClick = async e => {
        // Prevent spamming the submit button
        if(state.loading) return;

        setState(state => ({ ...state, loading: true }));

        try {
            if(!state.key) {
                setError({ message: "Encryption key required." });
                return;
            }
            const { data } = await encryptConfig({ key: state.key,
                                                   body: state.config });
            setOutput(data);
        } catch(e) {
            setError(e);
        } finally {
            setState(state => ({ ...state, loading: false }));
        }
    };

    return (
            <div className="h-100">
              <div className="form-group">
                <label htmlFor="key">Encryption Key</label>
                <input id="key"
                       type="text"
                       onInput={_handleKeyChange}
                       className="form-control mb-2" />
              </div>
              <AceEditor
                mode="json"
                theme="github"
                name="json-editor"
                value={ state.config }
                className="mb-2"
                onChange={ _handleConfigChange } />
              <button
                type="button"
                onClick={ _handleClick }
                className="btn btn-primary">
                { state.loading ? "Loading..." : "Encrypt" }
              </button>
            </div>
    );
};


export default Editor;
