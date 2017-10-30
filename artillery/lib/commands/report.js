/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

'use strict';

const fs = require('fs');
const path = require('path');
const l = require('lodash');
const openfile = require('open');

module.exports = report;

function report(jsonReportPath, options) {

  let reportFilename = options.output || jsonReportPath + '.html';

  let data = JSON.parse(fs.readFileSync(jsonReportPath, 'utf-8'));
  let templateFn = path.join(
    path.dirname(__filename),
    '../report/index.html.ejs');
  let template = fs.readFileSync(templateFn, 'utf-8');
  let compiledTemplate = l.template(template);
  let html = compiledTemplate({report: JSON.stringify(data, null, 2)});
  fs.writeFileSync(
    reportFilename,
    html,
    {encoding: 'utf-8', flag: 'w'});
  console.log('Report generated: %s', reportFilename);

  if (!options.output) {
    openfile(reportFilename);
  }
}