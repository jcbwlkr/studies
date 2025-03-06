const TaskService = {
  getCategories() {
    // Mimic calling an API that returns an array of objects.
    return Promise.resolve([
      {id: 1, name: "Chores"},
      {id: 2, name: "Work"},
      {id: 3, name: "Magic"},
    ])
  },
  getTasksByCategory(category) {
    switch(category.id) {
      case 1:
        return Promise.resolve([
          {id: 1, name: "Dishes"},
          {id: 2, name: "Laundry"},
          {id: 3, name: "Sweep"},
        ]);
      case 2:
        return Promise.resolve([
          {id: 4, name: "Chatbot quizzes"},
          {id: 5, name: "Put out fire"},
          {id: 6, name: "Prep for new hire"},
        ]);
      case 3:
        return Promise.resolve([
          {id: 7, name: "Wraiths deck"},
          {id: 8, name: "Marisi deck"},
          {id: 9, name: "Elsa deck"},
        ]);
      default:
        return Promise.resolve([]);
    }
  }
}
export {TaskService}
