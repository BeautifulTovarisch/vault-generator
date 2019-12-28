'use strict';

import React, { Component } from 'react';
import ReactDom from 'react-dom';

import Editor from './Editor/editor';

class App extends Component {
    render() {
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
                    <Editor />
                  </div>
                </div>
              </div>
            </div>
        );
    }
}

ReactDom.render(<App />, document.getElementById('react-mount-point'));
