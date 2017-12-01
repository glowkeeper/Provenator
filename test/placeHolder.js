var should = require('should');
('foobar').should.be.type('string');
// then that actual value '==' expected value
('foobar' == 'foobar').should.be.ok;
// then that actual value '===' expected value
('foobar').should.be.equal('foobar');
(5).should.be.exactly(5).and.be.a.Number();
