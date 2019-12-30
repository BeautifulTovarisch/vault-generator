'use strict';

import React from 'react';

const Output = ({ error, output }) =>
          <div className="h-100">
            { <p className="text-danger">{ error.message }</p> }
            <pre>{ output }</pre>
          </div>;


export default Output;
