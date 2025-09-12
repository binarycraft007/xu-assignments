import React from 'react';

function ClearButtons({ clearCompleted, clearAll }) {
  return (
    <div className="clear-buttons">
      <button onClick={clearCompleted}>Clear Completed</button>
      <button className="delete-btn" onClick={clearAll}>Clear All</button>
    </div>
  );
}

export default ClearButtons;
