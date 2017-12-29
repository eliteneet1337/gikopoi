/**
 * Entry-point for React Code in the Browser.
 */
import React from "react";
import ReactDOM from "react-dom";
import Layout from "./components/Layout";

window.PIXI = require('phaser/build/custom/pixi');
window.p2 = require('phaser/build/custom/p2');
window.Phaser = require('phaser/build/custom/phaser-split');
window.save = require('phaser-plugin-isometric/dist/phaser-plugin-isometric');

const app = document.getElementById('app');
ReactDOM.render(<Layout/>, app);
