/**
 * Parent component that switches between <LandingDisplay> and <GameDisplay>.
 */
import React from "react";
import LandingDisplay from "./LandingDisplay";
import GameDisplay from "./GameDisplay";
import ServerConnection from "../ServerConnection";


export default class Layout extends React.Component {

  constructor() {
    super();
    this.state = {initOpts: null, initResp: null, ws: null};
    this.onInitOpts = this.onInitOpts.bind(this);
    this.onInitResp = this.onInitResp.bind(this);
  }

  onInitOpts(initOpts) {
    // Generate the web socket connection
    const ws = new ServerConnection({onOpen: () => ws.clientInit(initOpts)});
    ws.on('clientInitResp', (initResp) => this.onInitResp(initResp));
    this.setState({ws});
  }

  onInitResp(initResp) {
    this.setState({initResp});
  }

  render() {
    if (this.state.initResp) {
      return <GameDisplay initResp={this.state.initResp} />
    } else {
      return <LandingDisplay onInitOpts={this.onInitOpts}/>
    }
  }
}
