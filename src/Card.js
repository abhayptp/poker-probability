import React, {Component} from 'react';
import CardSelector from './CardSelector.js';


export default class Cards extends Component {

  constructor(props) {
    super(props);
    this.state = {
      card: '',
      showCardSelector: false,
    }
    this.handler = this.handler.bind(this);
  }

  handler(c) {
    this.setState(c)
  }

  render() {
    var selectedCard = 'back_covers/blue.png'

    if(this.state.card) {
      selectedCard = 'cards/' + this.state.card;
    }

    return (
      <div>
        <div className='card' onClick={() => 
          this.setState({card: this.state.card, showCardSelector: true})}>

         <img src={'/assets/'+selectedCard} alt={selectedCard}/>
        </div>
        <div className='card-selector'>
          {this.state.showCardSelector ?
            <CardSelector />:
            null 
          }
        </div>
      </div>
    )
  }

}
