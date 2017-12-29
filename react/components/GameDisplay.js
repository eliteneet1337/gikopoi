/**
 * Manages the Game on the Page and runs the Game's Update Cycle.
 *
 * The game component holds a single <div /> for the game's display. After
 * rendering, the component loads the game onto the DOM with the Phaser
 * framework.
 */
import React from "react";
import settings from "../settings";


export default class GameDisplay extends React.Component {

  constructor(props) {
    super(props);

    this.game = null;
    this.char = null;
    this.cursors = null;

    this.preload = this.preload.bind(this);
    this.create = this.create.bind(this);
    this.update = this.update.bind(this);
    this.phaserRender = this.phaserRender.bind(this);
  }

  /**
   * Called after the render() method is complete and the DOM is built.
   */
  componentDidMount() {
    this.game = new Phaser.Game(800, 600, Phaser.AUTO, 'GameDisplay', {
      preload: this.preload, create: this.create, update: this.update,
      render: this.phaserRender
    });
  }

  /**
   * Game method of the Phaser API
   */
  preload() {
    // Load assets
    this.game.load.image('background', settings.ASSETS_URL + "/gikopoi_bitmaps/666.png");
    this.game.load.image('character', settings.ASSETS_URL + "/gikopoi_bitmaps/4.png");

    // Load and start isometric plugin
    this.game.plugins.add(new Phaser.Plugin.Isometric(this.game));
    this.game.time.advancedTiming = true;
    this.game.physics.startSystem(Phaser.Plugin.Isometric.ISOARCADE);
    //this.game.iso.anchor.setTo(0.5, 0.2);
  }

  /**
   * Game method of the Phaser API
   */
  create() {
    // Load sprites and enable physics
    this.bg = this.game.add.isoSprite(0, 0, 0, 'background');
    this.char = this.game.add.isoSprite(500, 380, 0, 'character');

    this.game.physics.isoArcade.enable(this.char);
    this.game.physics.isoArcade.gravity.setTo(0, 0, -500);
    this.char.body.collideWorldBounds = true
    this.game.world.setBounds(0, 0, 1920, 1920);
    this.game.camera.follow(this.char);

    // Setup keyboard interaction
    this.cursors = this.game.input.keyboard.createCursorKeys();

    this.game.input.keyboard.addKeyCapture([
        Phaser.Keyboard.LEFT,
        Phaser.Keyboard.RIGHT,
        Phaser.Keyboard.UP,
        Phaser.Keyboard.DOWN,
        Phaser.Keyboard.SPACEBAR
    ]);

    var space = this.game.input.keyboard.addKey(Phaser.Keyboard.SPACEBAR);

    space.onDown.add(function () {
        this.char.body.velocity.z = 300;
    }, this);

    this.char.body.velocity.y = 100;
  }

  /**
   * Game method of the Phaser API
   */
  update() {
    // Move the this.char at this speed.
    var speed = 100;

    if (this.cursors.up.isDown) {
        this.char.body.velocity.y = -speed;
    }
    else if (this.cursors.down.isDown) {
        this.char.body.velocity.y = speed;
    }
    else {
        this.char.body.velocity.y = 0;
    }

    if (this.cursors.left.isDown) {
        this.char.body.velocity.x = -speed;
    }
    else if (this.cursors.right.isDown) {
        this.char.body.velocity.x = speed;
    }
    else {
        this.char.body.velocity.x = 0;
    }
  }

  /**
   * Game method of the Phaser API
   */
  phaserRender() {
    this.game.debug.cameraInfo(this.game.camera, 32, 32);
    this.game.debug.spriteCoords(this.char, 32, 500);
  }

  /**
   * Returns the HTML to display for this component.
   */
  render() {
    const styles = {
      width: "800px",
      height: "600px",
      marginLeft: "auto",
      marginRight: "auto",
    }

    return <div id="GameDisplay" style={styles} />
  }
}
