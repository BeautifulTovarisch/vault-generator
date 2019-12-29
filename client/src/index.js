'use strict';

import React, { useState, useEffect, Component } from 'react';
import ReactDom from 'react-dom';

import "regenerator-runtime/runtime";

import Editor from './Editor/editor';

const App = () =>
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
        <Editor />
      </div>
      <div className="col-sm">
        <h2>Output</h2>
      </div>
    </div>
  </div>
</div>;

ReactDom.render(<App />, document.getElementById('react-mount-point'));
