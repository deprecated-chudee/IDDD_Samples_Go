package product

import (
	"iddd_samples_go/pkg/saasovation/agilepm/domain/discussion"
	"iddd_samples_go/pkg/saasovation/agilepm/domain/product/backlogitem"
	"iddd_samples_go/pkg/saasovation/agilepm/domain/team"
	"iddd_samples_go/pkg/saasovation/agilepm/domain/tenant"
)

type ID string

type Product struct {
	TenantID               tenant.ID
	ProductID              ID
	ProductOwnerID         team.ProductOwnerID
	Name                   string
	Description            string
	Discussion             Discussion
	DiscussionInitiationID string
	BacklogItems           []*backlogitem.BacklogItem
}

func NewProduct(tenantID tenant.ID, productID ID, productOwnerID team.ProductOwnerID, name string, description string, discussionAvailability discussion.Availability) (*Product, error) {
	discussion, err := fromAvailability(discussionAvailability)
	if err != nil {
		return nil, err
	}

	product := &Product{
		tenantID:       tenantID,
		productID:      productID,
		productOwnerID: productOwnerID,
		name:           name,
		description:    description,
		discussion:     *discussion,
		backlogItems:   make([]*backlogitem.BacklogItem, 0),
	}

	// TODO: DomainEventPublisher ProductCreated

	return product, nil
}

func (p *Product) ChangeProductOwner(productOwner team.ProductOwner) {
	if p.productOwnerID != productOwner.ID() {
		setProductOwnerID(productOwner.ID())

		// TODO: publish event
	}
}

func (p *Product) FailDiscussionInitiation() {
	if !p.discussion.availability.IsReady() {
		p.discussionInitiationID = nil
		//p.discussion = ProductDiscussion
		//.fromAvailability(
		//	DiscussionAvailability.FAILED))
	}
}

func (p *Product) InitiateDiscussion(descriptor *discussion.Descriptor) error {
	if err := p.AssertArgumentNotNil(descriptor,"The descriptor must not be null."); err != nil {
		return err
	}

	if p.discussion.availability.IsReady() {
		p.discussion = discussion.nowReady(descriptor)

		// TODO: publish event
		//DomainEventPublisher
		//.instance()
		//.publish(new ProductDiscussionInitiated(
		//	this.tenantId(),
		//	this.productId(),
		//	this.discussion()));
	}
}

func (p *Product) PlanBacklogItem(item backlogitem.BacklogItemID) {}

func (p *Product) PlannedProductBacklogItem(item backlogitem.BacklogItem) {}

func (p *Product) ReorderFrom(item backlogitem.BacklogItem) {}

func (p *Product) RequestDiscussion(item backlogitem.BacklogItem) {}

func (p *Product) ScheduleRelease(item backlogitem.BacklogItem) {}

func (p *Product) ScheduleSprint() {}

func (p *Product) StartDiscussionInitiation() {}