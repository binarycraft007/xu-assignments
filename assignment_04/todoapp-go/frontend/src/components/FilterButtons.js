import React from 'react';

function FilterButtons({ currentFilter, setFilter }) {
  return (
    <div className="filter-buttons">
      <button
        className={currentFilter === 'all' ? 'active' : ''}
        onClick={() => setFilter('all')}
      >
        All
      </button>
      <button
        className={currentFilter === 'unfinished' ? 'active' : ''}
        onClick={() => setFilter('unfinished')}
      >
        Unfinished
      </button>
      <button
        className={currentFilter === 'completed' ? 'active' : ''}
        onClick={() => setFilter('completed')}
      >
        Completed
      </button>
    </div>
  );
}

export default FilterButtons;
