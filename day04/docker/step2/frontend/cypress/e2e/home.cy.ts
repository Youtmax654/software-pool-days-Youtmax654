describe('Good home page', () => {
  it("Good title", () => {
    cy.visit('');
    cy.get('[data-cy=home-title]').should('contain', 'Nest Todolist powered by React');
  })
});
