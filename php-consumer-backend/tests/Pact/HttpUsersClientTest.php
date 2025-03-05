<?php

declare(strict_types=1);

namespace App\Tests\Infrastructure;

use PHPUnit\Framework\TestCase;
use App\Infrastructure\HttpUsersClient;
use PhpPact\Standalone\MockService\MockServerConfig;
use PhpPact\Consumer\InteractionBuilder;
use PhpPact\Consumer\Matcher\Matcher;
use PhpPact\Consumer\Model\ConsumerRequest;
use PhpPact\Consumer\Model\ProviderResponse;

final class HttpUsersClientTest extends TestCase
{
    private InteractionBuilder $pact;
    private MockServerConfig $config;

    protected function setUp(): void
    {
        $config = (new MockServerConfig())
            ->setLogLevel('debug')
            ->setConsumer('PHPBackendConsumer')
            ->setProvider('Backend')
            ->setPactDir(__DIR__.'/../../pacts');

        $this->config = $config;
        $this->pact = new InteractionBuilder($config);
    }

    public function testGetUserDetails()
    {
        $matcher = new Matcher();

        $this->pact
            ->given('A user with ID 1 exists')
            ->uponReceiving('A request for user details')
            ->with(
                (new ConsumerRequest())
                ->setMethod('GET')
                ->setPath('/users/1')
            )
            ->willRespondWith(
                (new ProviderResponse())
                ->setStatus(200)
                ->setHeaders(['Content-Type' => 'application/json'])
                ->setBody([
                    'id' => $matcher->like(1),
                    'email' => $matcher->like('john.doe@example.com'),
                ]),
            );

        $client = new HttpUsersClient((string) $this->config->getBaseUri());
        $userDetails = $client->getUserDetails(1);

        $this->assertEquals(1, $userDetails->id);
        $this->assertEquals('john.doe@example.com', $userDetails->email);

        $this->pact->verify();
    }
}
