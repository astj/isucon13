#! /usr/bin/env node
const { exec } = require("child_process");
const { stderr } = require("process");

const deploy = async () =>
  Promise.all(
    ["isucon1", "isucon2", "isucon3"].map(
      (s) =>
        new Promise((resolve, reject) => {
          console.log("deploy: ", s);
          exec(`ssh -t ${s} "sudo /home/isucon/webapp/deploy.sh"`, (e, stdout, steerr) => {
            if (e) {
              console.log('error!: ', stderr);
              reject(e);
            } else {
              console.log('success: ', stdout);
              resolve();
            }
          });
        })
    )
  );

(async () => {
  console.log("deploy");
  await deploy();
})();