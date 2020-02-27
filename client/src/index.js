'use strict';

import React, { useState, useEffect, Component } from 'react';
import ReactDom from 'react-dom';

import "regenerator-runtime/runtime";

import Editor from './Editor/editor';

const App = () => {
    const [state, setState] = useState({
        error: {},
        response: ""
    });

    const setError = err => setState(state => ({...state, error: err }));
    const setResponse = res => setState(state => ({...state, response: res}));

    useEffect(() => {
        if (!state.response || state.error.message) return;

        const link = document.createElement('a');
        link.href = URL.createObjectURL(new Blob([state.response]));
        link.download = "vault";

        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    }, [state.response]);

    return (
            <div id="wrapper" className="h-100">
              <div className="container">
                <div className="row">
                  <div className="text-center col-sm">
                    <h1 className="text-muted">Vault Generator</h1>
                    <p className="lead">Generate encrypted configuration file for use with Ansible Vault</p>
                  </div>
                </div>
                  <div className="row">
                    <div className="col-sm">
                      <h2>Input</h2>
                      { state.error && <p className="text-danger">{ state.error.message }</p> }
                      <Editor setError={setError} setResponse={setResponse} />
                    </div>
                  </div>
              </div>
            </div>
    );
};

ReactDom.render(<App />, document.getElementById('react-mount-point'));
