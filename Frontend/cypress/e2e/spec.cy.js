describe('Click filters', () => {
  it('Clicks all the filter buttons', () => {
    cy.visit('/')

    cy.get('mat-checkbox[style="display: block;"]').click({ multiple: true });
  })
})