describe('Open VReference', () => {
  it('Opens VReference', () => {
    cy.visit('/')

    cy.get('mat-checkbox[style="display: block;"]').click({ multiple: true });
  })
})