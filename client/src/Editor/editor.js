'use strict';

import React, { Component } from 'react';

import AceEditor from 'react-ace';

import "ace-builds/webpack-resolver";
import "ace-builds/src-noconflict/mode-json";
import "ace-builds/src-noconflict/theme-github.js"

class Editor extends Component {
    render() {
        return (
          <div>
            <AceEditor
              mode="json"
              theme="github"
              name="json-editor" />
          </div>
        );
    }
}

export default Editor;
