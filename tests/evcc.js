import fs from "fs";
import waitOn from "wait-on";
import axios from "axios";
import { exec, execSync } from "child_process";
import playwrightConfig from "../playwright.config";

const BASE_URL = playwrightConfig.use.baseURL;

const DB_PATH = "./robotui.db";
const BINARY = "./robotui";

export async function start(config, database) {
  await _clean();
  if (database) {
    await _restoreDatabase(database);
  }
  await _start(config);
}

export async function stop() {
  await _stop();
  await _clean();
}

export async function restart(config) {
  await _stop();
  await _start(config);
}

export async function cleanRestart(config) {
  await _stop();
  await _clean();
  await _start(config);
}

async function _restoreDatabase(database) {
  console.log("loading database", { database });
  execSync(`sqlite3 ${DB_PATH} < tests/${database}`);
}

async function _start(config) {
  console.log("starting robotui", { config });
  const instance = exec(`robotui_DATABASE_DSN=${DB_PATH} ${BINARY} --config tests/${config}`);
  instance.stdout.pipe(process.stdout);
  instance.on("exit", (code) => {
    if (code !== 0) {
      throw new Error("robotui terminated", code);
    }
  });
  await waitOn({ resources: [BASE_URL] });
}

async function _stop() {
  console.log("shutting down robotui");
  await axios.post(BASE_URL + "/api/shutdown");
}

async function _clean() {
  if (fs.existsSync(DB_PATH)) {
    console.log("delete database", DB_PATH);
    fs.unlinkSync(DB_PATH);
  }
}
