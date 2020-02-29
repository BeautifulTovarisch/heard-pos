'use strict';

import React, { useState, useEffect, Component } from 'react';
import ReactDom from 'react-dom';

import "regenerator-runtime/runtime";

const App = () => {

    return (
            <div id="wrapper" className="h-100">
              <div className="container">
                <div className="row">
                </div>
                  <div className="row">
                  </div>
              </div>
            </div>
    );
};

ReactDom.render(<App />, document.getElementById('react-mount-point'));
