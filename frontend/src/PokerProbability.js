import React, { useState } from 'react';
import api from './Api';

const PokerProbability = () => {
  const [probability, setProbability] = useState(null);

  const getProbability = async () => {
    try {
      const response = await api.post('/get_probability', {
        // replace with your actual data
        community_cards: {
          cards: [
            {rank: "A", suit: "S"},
            {rank: "K", suit: "S"},
            {rank: "Q", suit: "S"},
            {rank: "J", suit: "S"},
            {rank: "1", suit: "S"}
        ]
        },
        player_cards: [
          {
            name: "player1",
            cards: [
                {rank: "A", suit: "H"},
                {rank: "K", suit: "H"}
            ]
        },
        {
            name: "player2",
            cards: [
                {rank: "2", suit: "D"},
                {rank: "3", suit: "D"}
            ]
        }
        ],
        simulation_rounds: 1000,
        players_count: 2,
      });
      setProbability(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <button onClick={getProbability}>Get Poker Probability</button>
      {probability && (
        <p>
          Probability: {probability.player_results.map((player) => (
            <span key={player.name}>
              {player.name}: Win - {player.win_probability}, Tie - {player.tie_probability}
            </span>
          ))}
        </p>
      )}
    </div>
  );
};

export default PokerProbability;
