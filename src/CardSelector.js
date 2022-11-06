import React, {Component} from 'react';
import Popup from 'reactjs-popup';

export default class CardSelector extends Component {
  
  constructor(props) {
    super(props);

  }

  render() {
    

    return (
      <Popup trigger={<button> Trigger </button>}position="mid center">
      </Popup>
    )
  }
}
