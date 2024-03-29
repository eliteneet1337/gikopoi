/**
 * This component holds the default display when the page is opened.
 *
 * This is a page that holds the landing page, displaying the welcome banner and
 * character selection form.
 */
import React from "react";
import settings from "../settings";


export default class LandingDisplay extends React.Component {

  constructor() {
    super();
    this.state = {
      name: "",
      character: null,
    };
  }

  render() {
    const welcomeBannerStyles = {
      border: "1px solid black",
      backgroundColor: "#dfedd6",
      width: "484px",
      height: "251px",
      marginLeft: "auto",
      marginRight: "auto",
      marginBottom: "5px",
    }

    const selectBoxStyles = {
      border: "1px solid black",
      backgroundColor: "white",
      width: "450px",
      marginTop: "5px",
      marginLeft: "auto",
      marginRight: "auto",
      padding: "10px",
    }

    const charRows = settings.CHARACTERS.map((char) => {
      const rowStyle = {height: "60px"};
      const onClick = () => this.setState({character: char.characterName});
      var rowClasses = "list-group-item";

      if (this.state.character == char.characterName) {
        rowClasses += " active";
      }

      return  (
        <a
          key={char.characterName}
          href="#"
          class={rowClasses}
          style={rowStyle}
          onClick={onClick}>

          {char.visibleName}

          <div style={{float: "right", backgroundColor: "white", borderRadius: "5px", padding: "5px"}}>
            <img src={char.icon}  />
          </div>
        </a>
      )
    });

    const buttonStyles = {float: "right"};

    return (
      <div>
        <div style={welcomeBannerStyles} align>
          <img src="static/welcome_banner.png" />
        </div>

        <div style={selectBoxStyles}>
          <div class="input-group">
            <span class="input-group-addon" id="basic-addon1">
              <span class="glyphicon glyphicon-user" aria-hidden="true"></span>
            </span>

            <input
              type="text"
              class="form-control"
              placeholder="name"
              value={this.state.value}
              onChange={(e) => this.setState({name: e.target.value})}
              aria-describedby="basic-addon1" />
          </div>

          <br />

          <div class="list-group">
            {charRows}
          </div>

          <br />

          <a
            href="#"
            class="btn btn-primary"
            role="button"
            style={buttonStyles}
            onClick={() => this.props.onInitOpts(this.state)}>
            LOGIN
          </a>

          <br />
          <br />
        </div>
      </div>
    )
  }
}
